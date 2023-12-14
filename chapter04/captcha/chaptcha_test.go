package captcha_test

import (
	"go-swe/chapter04/captcha"
	"image"
	"testing"
)

func TestChallangeUserSuccess(t *testing.T) {

	got := captcha.ChallengeUser(stubChallenger("42"), stubPrompter("42"))
	if got != true {
		t.Fatal("expected challangeUser to return true")
	}

}

func TestChallangeUserFail(t *testing.T) {

	got := captcha.ChallengeUser(stubChallenger("lorem ipsum"), stubPrompter("42"))
	if got != false {
		t.Fatal("expected challangeUser to return false")
	}

}

type stubChallenger string

func (c stubChallenger) Challenge() (image.Image, string) {
	return image.NewRGBA(image.Rect(0, 0, 100, 100)), string(c)
}

type stubPrompter string

func (p stubPrompter) Prompt(_ image.Image) string {
	return string(p)
}
