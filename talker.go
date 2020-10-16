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
	detect := WCharToString(C.Tolk_DetectScreenReader())
	return detect
}

func Output(text string, interrupt bool) (bool, error) {
	wcharstr, err := StringToWchart(text)
	if err != nil {
		return false, err
		}
	ok := C.Tolk_Output(wcharstr, C.bool(interrupt))
	return bool(ok), nil
}

