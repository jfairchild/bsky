package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/api/bsky"
	lexutil "github.com/bluesky-social/indigo/lex/util"
	"github.com/bluesky-social/indigo/util/cliutil"
	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
	cli "github.com/urfave/cli/v2"
)

// PrintISODateTime prints the current time in ISO 8601 format (RFC 3339)
func PrintISODateTime() string {
	now := time.Now().UTC()
	formattedTime := now.Format(time.RFC3339)
	return formattedTime
}

// PostToBsky posts a message to a bsky.social account
func PostToBsky(ctx context.Context, baseURL, privateKey, message string) error {

	myPost := &bsky.FeedPost{
		LexiconTypeID: "app.bsky.feed.post",
		Text:          message,
		CreatedAt:     PrintISODateTime(),
		Langs:         []string{"en-US"},
	}
	cctx := &cli.Context{}
	// Create a new bsky client
	xrpcc, err := cliutil.GetXrpcClient(cctx, true)
	if err != nil {
		return fmt.Errorf("failed to create bsky client: %w", err)
	}

	l := logr.FromContextOrDiscard(ctx)
	l.Info("Xrpc Client Auth...")
	auth := xrpcc.Auth

	resp, err := atproto.RepoCreateRecord(context.TODO(), xrpcc, &atproto.RepoCreateRecord_Input{
		Collection: "app.bsky.feed.post",
		Repo:       auth.Did,
		Record:     &lexutil.LexiconTypeDecoder{Val: myPost},
	})
	if err != nil {
		return fmt.Errorf("failed to create post: %w", err)
	}

	l.Info(resp.Cid)
	l.Info(resp.Uri)
	l.Info("Successfully posted message to bsky.social!")
	return nil
}

func main() {
	// Replace with your bsky.social base URL, private key and message
	baseURL := "https://bsky.social"
	privateKey := "YOUR_PRIVATE_KEY"
	message := "This is a test message from Go!"
	logger := stdr.New(log.New(os.Stderr, "", log.Default().Flags()))
	logger.Info("Welcome to bsky post", "bsky-url", baseURL)

	ctx := logr.NewContext(context.Background(), logger)
	// Set up OpenTelemetry.
	otelShutdown, err := setupOTelSDK(ctx)
	if err != nil {
		return
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	err = PostToBsky(ctx, baseURL, privateKey, message)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
