package talker


//#include <wchar.h>
//#include <stdbool.h>
//#include <fcntl.h>
//#include "Tolk.h"
import "C"

	import "golang.org/x/sys/windows"

func WCharToString(p *C.wchar_t) string {
	return windows.UTF16PtrToString((*uint16)(p))
}

func StringToWchart(s string) (*C.wchar_t, error) {
	p, err := windows.UTF16PtrFromString(s)
	return (*C.wchar_t)(p), err
}

func Load() {
C.Tolk_Load()
}

func Unload() {
	C.Tolk_Unload()
}

func IsLoaded() bool {
	return bool(C.Tolk_IsLoaded())
}

func TrySAPI(yesno bool) {
	C.Tolk_TrySAPI(C.bool(yesno))
}

func PreferSAPI(yesno bool) {
	C.Tolk_PreferSAPI(C.bool(yesno))
}

func DetectScreenReader() string {
	return WCharToString(C.Tolk_DetectScreenReader())
}

func Output(text string, interrupt bool) (bool, error) {
	wcharstr, err := StringToWchart(text)
	if err != nil {
		return false, err
		}
	return bool(C.Tolk_Output(wcharstr, C.bool(interrupt))), nil
}

// C Description:  Tests if the current screen reader driver supports speech output, if one is set. If none is set, tries to detect the currently active screen reader before testing for speech support. You should call Tolk_Load once before using this function.
func HasSpeech() bool {
	return bool(C.Tolk_HasSpeech())
}

// C Description:  Tests if the current screen reader driver supports braille output, if one is set. If none is set, tries to detect the currently active screen reader before testing for braille support. You should call Tolk_Load once before using this function.
func HasBraille() bool {
	return bool(C.Tolk_HasBraille())
}

// C Description:  Speaks text through the current screen reader driver, if one is set and supports speech output. If none is set or if it encountered an error, tries to detect the currently active screen reader before speaking the text. Use this function only if you specifically need to speak text through the current screen reader without also brailling it. Not all screen reader drivers may support this functionality. Therefore, use Tolk_Output whenever possible. You should call Tolk_Load once before using this function. This function is asynchronous.
func Speak(value string, interrupt bool) bool {
	v, err := StringToWchart(value)
	if err != nil {
		return false
		}
	return bool(C.Tolk_Speak(v, C.bool(interrupt)))
}

// C Description:  Brailles text through the current screen reader driver, if one is set and supports braille output. If none is set or if it encountered an error, tries to detect the currently active screen reader before brailling the given text. Use this function only if you specifically need to braille text through the current screen reader without also speaking it. Not all screen reader drivers may support this functionality. Therefore, use Tolk_Output whenever possible. You should call Tolk_Load once before using this function.
func Braille(value string) bool {
	v, err := StringToWchart(value)
	if err != nil {
		return false
		}
	return bool(C.Tolk_Braille(v))
}

// C Description:  Tests if the screen reader associated with the current screen reader driver is speaking, if one is set and supports querying for status information. If none is set, tries to detect the currently active screen reader before testing if it is speaking. You should call Tolk_Load once before using this function.
func IsSpeaking() bool {
	return bool(C.Tolk_IsSpeaking())
}

// C Description:  Silences the screen reader associated with the current screen reader driver, if one is set and supports speech output. If none is set or if it encountered an error, tries to detect the currently active screen reader before silencing it. You should call Tolk_Load once before using this function.
func Silence() bool {
	return bool(C.Tolk_Silence())
}
