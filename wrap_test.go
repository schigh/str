package str

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const WrapTestLiteral = `TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4gQWVuZWFuIGxhY2luaWEgcXVpcyBxdWFtIGV0IHBvc3VlcmUuIEluIGZhY2lsaXNpcywgdG9ydG9yIGlkIGNvbmd1ZSBwaGFyZXRyYSwgbGVvIG9kaW8gbWF4aW11cyBsaWJlcm8sIGVnZXQgZ3JhdmlkYSBlbGl0IG51bmMgZXQgbmlzbC4gRnVzY2UgZWdldCB1cm5hIHZlbCBsaWd1bGEgc29sbGljaXR1ZGluIHBvc3VlcmUgaWQgdml0YWUgcHVydXMuIFBlbGxlbnRlc3F1ZSBoYWJpdGFudCBtb3JiaSB0cmlzdGlxdWUgc2VuZWN0dXMgZXQgbmV0dXMgZXQgbWFsZXN1YWRhIGZhbWVzIGFjIHR1cnBpcyBlZ2VzdGFzLiBWZXN0aWJ1bHVtIHF1aXMgbGliZXJvIGRpY3R1bSwgY29uZGltZW50dW0gdHVycGlzIGlkLCB0ZW1wb3IgbWV0dXMuIEluIHNlZCBuaWJoIHF1YW0uIFZlc3RpYnVsdW0gZXUgYXVndWUgbmliaC4gSW4gbWFnbmEgbGVjdHVzLCBmYXVjaWJ1cyBldCBlbmltIG5lYywgbGFvcmVldCBwb3N1ZXJlIG51bmMuIE51bGxhIHRlbXBvciwgZW5pbSBldCB1bHRyaWNlcyBldWlzbW9kLCBqdXN0byBleCB0aW5jaWR1bnQgYXVndWUsIG5lYyBhdWN0b3IgbWV0dXMgbG9yZW0gaW4gZWxpdC4=`
const WrapResultLiteral = `TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4g
QWVuZWFuIGxhY2luaWEgcXVpcyBxdWFtIGV0IHBvc3VlcmUuIEluIGZhY2lsaXNpcywgdG9ydG9y
IGlkIGNvbmd1ZSBwaGFyZXRyYSwgbGVvIG9kaW8gbWF4aW11cyBsaWJlcm8sIGVnZXQgZ3Jhdmlk
YSBlbGl0IG51bmMgZXQgbmlzbC4gRnVzY2UgZWdldCB1cm5hIHZlbCBsaWd1bGEgc29sbGljaXR1
ZGluIHBvc3VlcmUgaWQgdml0YWUgcHVydXMuIFBlbGxlbnRlc3F1ZSBoYWJpdGFudCBtb3JiaSB0
cmlzdGlxdWUgc2VuZWN0dXMgZXQgbmV0dXMgZXQgbWFsZXN1YWRhIGZhbWVzIGFjIHR1cnBpcyBl
Z2VzdGFzLiBWZXN0aWJ1bHVtIHF1aXMgbGliZXJvIGRpY3R1bSwgY29uZGltZW50dW0gdHVycGlz
IGlkLCB0ZW1wb3IgbWV0dXMuIEluIHNlZCBuaWJoIHF1YW0uIFZlc3RpYnVsdW0gZXUgYXVndWUg
bmliaC4gSW4gbWFnbmEgbGVjdHVzLCBmYXVjaWJ1cyBldCBlbmltIG5lYywgbGFvcmVldCBwb3N1
ZXJlIG51bmMuIE51bGxhIHRlbXBvciwgZW5pbSBldCB1bHRyaWNlcyBldWlzbW9kLCBqdXN0byBl
eCB0aW5jaWR1bnQgYXVndWUsIG5lYyBhdWN0b3IgbWV0dXMgbG9yZW0gaW4gZWxpdC4=`
const WrapTestBeforeAfter = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean lacinia quis quam et posuere. In facilisis, tortor id congue pharetra, leo odio maximus libero, eget gravida elit nunc et nisl. Fusce eget urna vel ligula sollicitudin posuere id vitae purus. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum quis libero dictum, condimentum turpis id, tempor metus. In sed nibh quam. Vestibulum eu augue nibh.`
const WrapResultBeforeWord = `Lorem ipsum dolor sit amet,
consectetur adipiscing elit.
Aenean lacinia quis quam et
posuere. In facilisis, tortor id
congue pharetra, leo odio
maximus libero, eget gravida
elit nunc et nisl. Fusce eget
urna vel ligula sollicitudin
posuere id vitae purus.
Pellentesque habitant morbi
tristique senectus et netus et
malesuada fames ac turpis
egestas. Vestibulum quis libero
dictum, condimentum turpis id,
tempor metus. In sed nibh quam.
Vestibulum eu augue nibh.`
const WrapResultAfterWord = `Lorem ipsum dolor sit amet, consectetur
adipiscing elit. Aenean lacinia
quis quam et posuere. In facilisis,
tortor id congue pharetra, leo odio
maximus libero, eget gravida elit
nunc et nisl. Fusce eget urna vel
ligula sollicitudin posuere id vitae
purus. Pellentesque habitant morbi
tristique senectus et netus et malesuada
fames ac turpis egestas. Vestibulum
quis libero dictum, condimentum
turpis id, tempor metus. In sed
nibh quam. Vestibulum eu augue nibh.`

func TestWrapLiteral(t *testing.T) {
	options := &WrapOptions{
		Width:     76,
		LineBreak: "\n",
		Behavior:  WrapLiteral,
	}

	out := WrapWithOptions(WrapTestLiteral, options)
	assert.Equal(t, WrapResultLiteral, out)
}

func TestWrapBeforeWord(t *testing.T) {
	options := &WrapOptions{
		Width:     32,
		LineBreak: "\n",
		Behavior:  WrapBeforeWord,
	}
	out := WrapWithOptions(WrapTestBeforeAfter, options)
	assert.Equal(t, WrapResultBeforeWord, out)
}

func TestWrapAfterWord(t *testing.T) {
	options := &WrapOptions{
		Width:     32,
		LineBreak: "\n",
		Behavior:  WrapAfterWord,
	}
	out := WrapWithOptions(WrapTestBeforeAfter, options)
	assert.Equal(t, WrapResultAfterWord, out)
}
