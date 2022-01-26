//go:build windows || !linux
// +build windows !linux

// Copyright (C) 2020, 2021, 2022 SeanTolstoyevski - mailto:seantolstoyevski@protonmail.com
//
// The source code of this project is licensed under the MIT license.
// But this package has dependencies licensed with different licenses.
// You can find the license on the repo's main folder.
// Provided without warranty of any kind.

package talker

import (
	"syscall"
	"unsafe"
)

var (
	tolkLib = syscall.NewLazyDLL("Tolk.dll")

	procLoadTolk = tolkLib.NewProc("Tolk_Load")

	procUnloadTolk = tolkLib.NewProc("Tolk_Unload")

	procIsLoadedTolk = tolkLib.NewProc("Tolk_IsLoaded")

	procTrySAPITolk = tolkLib.NewProc("Tolk_TrySAPI")

	procPreferSAPITolk = tolkLib.NewProc("Tolk_PreferSAPI")

	procDetectScreenReaderTolk = tolkLib.NewProc("Tolk_DetectScreenReader")

	procOutputTolk = tolkLib.NewProc("Tolk_Output")

	procHasSpeechTolk = tolkLib.NewProc("Tolk_HasSpeech")

	procHasBrailleTolk = tolkLib.NewProc("Tolk_HasBraille")

	procSpeakProcTolk = tolkLib.NewProc("Tolk_Speak")

	procBrailleTolk = tolkLib.NewProc("Tolk_Braille")

	procIsSpeakingTolk = tolkLib.NewProc("Tolk_IsSpeaking")

	procSilenceTolk = tolkLib.NewProc("Tolk_Silence")
)

// syscall.StringToUTF16Ptr(

// Load load Talker. When you're done, you should call Unload.
//
// C Description: Initializes Tolk by loading and initializing the screen reader drivers and setting the current screen reader driver, provided at least one of the supported screen readers is active.
// Also initializes COM if it has not already been initialized on the calling thread.
// Calling this function more than once will only initialize COM. You should call this function before using the functions below.
//
// Use IsLoaded to determine if Tolk has been initialized.
func Load() {
	procLoadTolk.Call()
}

// Unload release Talker's resource. You should call it, if finished  you're transactions.
//
// Like exiting the program. You shouldn't use it with  expression like Defer. Because you have to load (Load()) it again.
//
// C Description:  Finalizes Tolk by finalizing and unloading the screen reader drivers and clearing the current screen reader driver, provided one was set. Also uninitializes COM on the calling thread. Calling this function more than once will only uninitialize COM. You should not use the functions below if this function has been called.
func Unload() {
	procUnloadTolk.Call()
}

// IsLoaded returns whether Talker is loaded. True if loaded, false otherwise.
//
// C Description: Tests if Tolk has been initialized.
func IsLoaded() bool {
	r, _, _ := procIsLoadedTolk.Call()
	return r != 0
}

// If you want to use SAPI it might be better to call it before Load().
//
// C Description:  Sets if Microsoft Speech API (SAPI) should be used in the screen reader auto-detection process. The default is not to include SAPI. The SAPI driver will use the system default synthesizer, voice and soundcard. This function triggers the screen reader detection process if needed.
// For best performance, you should call this function before calling Load.
func TrySAPI(yesno bool) {
	procTrySAPITolk.Call(uintptr(unsafe.Pointer(&yesno)))
}

// C Description:  If auto-detection for SAPI has been turned on through Tolk_TrySAPI, sets if SAPI should be placed first (true) or last (false) in the screen reader detection list. Putting it last is the default and is good for using SAPI as a fallback option. Putting it first is good for ensuring SAPI is used even when a screen reader is running, but keep in mind screen readers will still be tried if SAPI is unavailable. This function triggers the screen reader detection process if needed.
//
// For best performance, you should call this function before calling Load.
func PreferSAPI(yesno bool) {
	procPreferSAPITolk.Call(uintptr(unsafe.Pointer(&yesno)))
}

// DetectScreenReader returns name of current screen reader. Like NVDA, JAWS...
//
// C Description:  Returns the common name for the currently active screen reader driver, if one is set.
// If none is set, tries to detect the currently active screen reader before looking up the name.
// If no screen reader is active, NULL is returned.
// Note that the drivers hard-code the common name, it is not requested from the screen reader itself.
// You should call Load once before using this function.
func DetectScreenReader() string {
	r, _, _ := procDetectScreenReaderTolk.Call()
	v := *(*[]uint16)(unsafe.Pointer(&r))
	return syscall.UTF16ToString(v)
}

// Output sends text to be speaking  to the current screen reader.
// If "interrupt" is True, the current speak is interrupt.
//
// Users generally do not like to interrupt their existing speaking.
//
// C Description:  Outputs text through the current screen reader driver, if one is set.
// If none is set or if it encountered an error, tries to detect the currently active screen reader before outputting the text.
// This is the preferred function to use for sending text to a screen reader, because it uses all of the supported output methods (speech and/or braille depending on the current screen reader driver).
//
// You should call Load once before using this function.
//
// This function is asynchronous.
func Output(text string, interrupt bool) bool {
	r, _, _ := procOutputTolk.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))), uintptr(unsafe.Pointer(&interrupt)))
	return r != 0
}

// C Description:  Tests if the current screen reader driver supports speech output, if one is set.
// If none is set, tries to detect the currently active screen reader before testing for speech support.
//
// You should call Load once before using this function.
func HasSpeech() bool {
	r, _, _ := procHasSpeechTolk.Call()
	return r != 0
}

// C Description:  Tests if the current screen reader driver supports braille output, if one is set.
// If none is set, tries to detect the currently active screen reader before testing for braille support.
//
// You should call Load once before using this function.
func HasBraille() bool {
	r, _, _ := procHasBrailleTolk.Call()
	return r != 0
}

// C Description:  Speaks text through the current screen reader driver, if one is set and supports speech output.
// If none is set or if it encountered an error, tries to detect the currently active screen reader before speaking the text. Use this function only if you specifically need to speak text through the current screen reader without also brailling it. Not all screen reader drivers may support this functionality. Therefore, use Tolk_Output whenever possible.
//
// You should call Load once before using this function. This function is asynchronous.
func Speak(value string, interrupt bool) bool {
	r, _, _ := procSpeakProcTolk.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(value))), uintptr(unsafe.Pointer(&interrupt)))
	return r != 0
}

// C Description:  Brailles text through the current screen reader driver, if one is set and supports braille output.
// If none is set or if it encountered an error, tries to detect the currently active screen reader before brailling the given text. Use this function only if you specifically need to braille text through the current screen reader without also speaking it. Not all screen reader drivers may support this functionality. Therefore, use Tolk_Output whenever possible.
//
// You should call Load once before using this function.
func Braille(text string) bool {
	r, _, _ := procBrailleTolk.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))))
	return r != 0
}

// C Description:  Tests if the screen reader associated with the current screen reader driver is speaking, if one is set and supports querying for status information.
// If none is set, tries to detect the currently active screen reader before testing if it is speaking.
//
// You should call Load once before using this function.
func IsSpeaking() bool {
	r, _, _ := procIsSpeakingTolk.Call()
	return r != 0
}

// C Description:  Silences the screen reader associated with the current screen reader driver, if one is set and supports speech output.
// If none is set or if it encountered an error, tries to detect the currently active screen reader before silencing it.
//
// You should call Load once before using this function.
func Silence() bool {
	r, _, _ := procSilenceTolk.Call()
	return r != 0
}
