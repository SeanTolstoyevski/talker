//Talker is a wrapper for the Tolk library.  
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

//#include <wchar.h>
//#include <stdbool.h>
//#include <fcntl.h>
//#include "Tolk.h"
import "C"

	import "golang.org/x/sys/windows"

// Convert UTF16 (wchar_t) c to GoString.  
//  
//  You don't need to use it. Talker convert the strings you send to suitable formats.
func WCharToString(p *C.wchar_t) string {
	return windows.UTF16PtrToString((*uint16)(p))
}

// Convert go string  c to wchar_t.  
//  
// You don't need to use it.
func StringToWchart(s string) (*C.wchar_t, error) {
	p, err := windows.UTF16PtrFromString(s)
	return (*C.wchar_t)(p), err
}

// Load Talker. When you're done, you should call Unload.  
//  
// C Description: Initializes Tolk by loading and initializing the screen reader drivers and setting the current screen reader driver, provided at least one of the supported screen readers is active. Also initializes COM if it has not already been initialized on the calling thread. Calling this function more than once will only initialize COM. You should call this function before using the functions below.  
//  
// Use IsLoaded to determine if Tolk has been initialized.
func Load() {
C.Tolk_Load()
}

// Release Talker's resource. You should call it, if finished  you're transactions.  
//  
// Like exiting the program. You shouldn't use it with  expression like Defer. Because you have to load (Load()) it again.
//  
// C Description:  Finalizes Tolk by finalizing and unloading the screen reader drivers and clearing the current screen reader driver, provided one was set. Also uninitializes COM on the calling thread. Calling this function more than once will only uninitialize COM. You should not use the functions below if this function has been called.
func Unload() {
	C.Tolk_Unload()
}

//Returns whether Talker is loaded. True if loaded, false otherwise. 
//  
// C Description: Tests if Tolk has been initialized.
func IsLoaded() bool {
	return bool(C.Tolk_IsLoaded())
}

// If you want to use SAPI it might be better to call it before Load().  
//  
// C Description:  Sets if Microsoft Speech API (SAPI) should be used in the screen reader auto-detection process. The default is not to include SAPI. The SAPI driver will use the system default synthesizer, voice and soundcard. This function triggers the screen reader detection process if needed.  
// For best performance, you should call this function before calling Load.
func TrySAPI(yesno bool) {
	C.Tolk_TrySAPI(C.bool(yesno))
}

// C Description:  If auto-detection for SAPI has been turned on through Tolk_TrySAPI, sets if SAPI should be placed first (true) or last (false) in the screen reader detection list. Putting it last is the default and is good for using SAPI as a fallback option. Putting it first is good for ensuring SAPI is used even when a screen reader is running, but keep in mind screen readers will still be tried if SAPI is unavailable. This function triggers the screen reader detection process if needed.  
//  
// For best performance, you should call this function before calling Load.
func PreferSAPI(yesno bool) {
	C.Tolk_PreferSAPI(C.bool(yesno))
}

// Returns name of current screen reader. Like NVDA, JAWS...
//  
// C Description:  Returns the common name for the currently active screen reader driver, if one is set. If none is set, tries to detect the currently active screen reader before looking up the name. If no screen reader is active, NULL is returned. Note that the drivers hard-code the common name, it is not requested from the screen reader itself.  
// You should call Load once before using this function.
func DetectScreenReader() string {
	return WCharToString(C.Tolk_DetectScreenReader())
}

// Sends text to be speaking  to the current screen reader. If "interrupt" is True, the current speak is interrupt.
//  
// Users generally do not like to interrupt their existing speaking.
//  
// C Description:  Outputs text through the current screen reader driver, if one is set. If none is set or if it encountered an error, tries to detect the currently active screen reader before outputting the text. This is the preferred function to use for sending text to a screen reader, because it uses all of the supported output methods (speech and/or braille depending on the current screen reader driver).  
//  
// You should call Load once before using this function.  
//  
// This function is asynchronous.
func Output(text string, interrupt bool) (bool, error) {
	wcharstr, err := StringToWchart(text)
	if err != nil {
		return false, err
		}
	return bool(C.Tolk_Output(wcharstr, C.bool(interrupt))), nil
}

// C Description:  Tests if the current screen reader driver supports speech output, if one is set. If none is set, tries to detect the currently active screen reader before testing for speech support.  
//  
// You should call Load once before using this function.
func HasSpeech() bool {
	return bool(C.Tolk_HasSpeech())
}

// C Description:  Tests if the current screen reader driver supports braille output, if one is set. If none is set, tries to detect the currently active screen reader before testing for braille support.  
//  
// You should call Load once before using this function.
func HasBraille() bool {
	return bool(C.Tolk_HasBraille())
}

// C Description:  Speaks text through the current screen reader driver, if one is set and supports speech output. If none is set or if it encountered an error, tries to detect the currently active screen reader before speaking the text. Use this function only if you specifically need to speak text through the current screen reader without also brailling it. Not all screen reader drivers may support this functionality. Therefore, use Tolk_Output whenever possible.  
//  
// You should call Load once before using this function. This function is asynchronous.
func Speak(value string, interrupt bool) bool {
	v, err := StringToWchart(value)
	if err != nil {
		return false
		}
	return bool(C.Tolk_Speak(v, C.bool(interrupt)))
}

// C Description:  Brailles text through the current screen reader driver, if one is set and supports braille output. If none is set or if it encountered an error, tries to detect the currently active screen reader before brailling the given text. Use this function only if you specifically need to braille text through the current screen reader without also speaking it. Not all screen reader drivers may support this functionality. Therefore, use Tolk_Output whenever possible.  
//  
// You should call Load once before using this function.
func Braille(value string) bool {
	v, err := StringToWchart(value)
	if err != nil {
		return false
		}
	return bool(C.Tolk_Braille(v))
}

// C Description:  Tests if the screen reader associated with the current screen reader driver is speaking, if one is set and supports querying for status information. If none is set, tries to detect the currently active screen reader before testing if it is speaking.  
//  
// You should call Load once before using this function.
func IsSpeaking() bool {
	return bool(C.Tolk_IsSpeaking())
}

// C Description:  Silences the screen reader associated with the current screen reader driver, if one is set and supports speech output. If none is set or if it encountered an error, tries to detect the currently active screen reader before silencing it.  
//  
// You should call Load once before using this function.
func Silence() bool {
	return bool(C.Tolk_Silence())
}
