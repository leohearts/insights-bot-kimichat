package smr

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/go-shiori/go-readability"
	"github.com/imroc/req/v3"
	"github.com/samber/lo"
	goopenai "github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"

	"github.com/nekomeowww/insights-bot/pkg/logger"
	"github.com/nekomeowww/insights-bot/pkg/openai"
)

type NewModelParams struct {
	fx.In

	OpenAIClient *openai.Client
	Logger       *logger.Logger
}

type Model struct {
	openai *openai.Client
	logger *logger.Logger
	req    *req.Client
}

func NewModel() func(NewModelParams) *Model {
	return func(param NewModelParams) *Model {
		return &Model{
			req:    req.C(),
			logger: param.Logger,
			openai: param.OpenAIClient,
		}
	}
}

var (
	ErrNetworkError        = errors.New("network error")
	ErrRequestFailed       = errors.New("request failed")
	ErrContentNotSupported = errors.New("content not supported")
)

type URLSummarizationOutput struct {
	URL   string
	Title string
	Msg   string
}

func (u *URLSummarizationOutput) FormatSummarizationAsHTML() string {
	return fmt.Sprintf("<b><a href=\"%s\">%s</a></b>\n%s\n\n<em>🤖️ Generated by chatGPT</em>", u.URL, u.Title, u.Msg)
}

// FormatSummarizationAsSlackMarkdown the link syntax in slack markdown flavor is different than standard
func (u *URLSummarizationOutput) FormatSummarizationAsSlackMarkdown() string {
	return fmt.Sprintf("*<%s|%s>*\n%s\n\n_🤖️ Generated by chatGPT_", u.URL, u.Title, u.Msg)
}

func (m *Model) SummarizeInputURL(url string) (*URLSummarizationOutput, error) {
	article, err := m.extractContentFromURL(url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse %s, %w", url, err)
	}

	textContent := m.openai.TruncateContentBasedOnTokens(article.TextContent, 3000)
	m.logger.WithFields(logrus.Fields{
		"title": article.Title,
		"url":   url,
	}).Infof("✍️ summarizing article...")
	resp, err := m.openai.SummarizeWithQuestionsAsSimplifiedChinese(
		context.Background(),
		article.Title,
		article.Byline,
		textContent,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create chat completion for summarizing... %w", err)
	}

	respMessages := lo.Map(resp.Choices, func(item goopenai.ChatCompletionChoice, _ int) string {
		return item.Message.Content
	})
	if len(respMessages) == 0 {
		return nil, fmt.Errorf("no response from OpenAI")
	}

	m.logger.WithFields(logrus.Fields{
		"title":                  article.Title,
		"url":                    url,
		"prompt_token_usage":     resp.Usage.PromptTokens,
		"completion_token_usage": resp.Usage.CompletionTokens,
		"total_token_usage":      resp.Usage.TotalTokens,
	}).Infof("✅ summarizing article done")
	return &URLSummarizationOutput{
		URL:   url,
		Title: article.Title,
		Msg:   respMessages[0],
	}, nil
}

func (m *Model) extractContentFromURL(urlString string) (*readability.Article, error) {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	if parsedURL == nil {
		return nil, errors.New("empty url")
	}

	resp, err := m.req.
		SetUserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 Edg/111.0.1661.54").
		SetTimeout(time.Minute).
		R().
		EnableDump().
		Get(parsedURL.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get url %s, %w: %v", parsedURL.String(), ErrNetworkError, err)
	}
	if !resp.IsSuccessState() {
		return nil, fmt.Errorf("failed to get url %s, %w, status code: %d, dump: %s", parsedURL.String(), ErrRequestFailed, resp.StatusCode, resp.Dump())
	}
	if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		return nil, fmt.Errorf("url fetched, but content-type not supported yet, %w, content-type: %s", ErrContentNotSupported, resp.Header.Get("Content-Type"))
	}

	buffer := new(bytes.Buffer)
	_, err = io.Copy(buffer, resp.Body)
	if err != nil {
		return nil, err
	}

	urlContent, err := readability.FromReader(buffer, parsedURL)
	if err != nil {
		return nil, err
	}

	return &urlContent, nil
}
