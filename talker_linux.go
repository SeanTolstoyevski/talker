//go:build linux || !windows
// +build linux !windows

// Copyright (C) 2020, 2021 SeanTolstoyevski - mailto:seantolstoyevski@protonmail.com
//
// The source code of this project is licensed under the MIT license.
// But this package has dependencies licensed with different licenses.
// You can find the license on the repo's main folder.
// Provided without warranty of any kind.

// Package talker is a wrapper for the Tolk library.
//
// It processes the API of many screen readers in a single API.
// You don't need to be particularly concerned with any low-level functions.
//
// Currently Tolk/Talker supports these screen readers: NVDA, JAWS, WindowEyes, SuperNova, SystemAccess, ZoomText.
//
// Microsoft Windows SAPI is also supported.
//
// Naming note: "Tolk_" part is taken from the names of functions.
// If you have used Tolk, you can adapt to Talker by simply deleting these letters.
//
// For example: C: Tolk_Load = Go: talker.Load()
//
// C: Tolk_Output = Go: talker.Output("Hello, I am Golang and CGO or Talker :)", false)
//
// Most of the project's DLL dependencies are licensed under the GNU license.
//
// So I can't statically linking  them.
//
// You must copy the appropriate DLLs for your architecture from x64 or x86 to your project directory. And license files (LICENSE-NVDA.txt, LICENSE_tolk.txt) :)
package talker

// DetectScreenReader
func DetectScreenReader() string {
	return ""
}

// Output sends text to be speaking  to the current screen reader. If "interrupt" is True, the current speak is interrupt.
func Output(text string, interrupt bool) (bool, error) {
	return false, nil
}

func HasSpeech() bool {
	return false
}

func HasBraille() bool {
	return false
}

func Speak(value string, interrupt bool) bool {
	return false
}

func Braille(value string) bool {
	return false
}

func IsSpeaking() bool {
	return false
}

func Silence() bool {
	return false
}
