// Copyright (C) 2020, 2021 SeanTolstoyevski - mailto:seantolstoyevski@protonmail.com
//
// The source code of this project is licensed under the MIT license.
// But this package has dependencies licensed with different licenses.
// You can find the license on the repo's main folder.
// Provided without warranty of any kind.

package talker

//#cgo windows CFLAGS: -Iinclude
//#cgo windows LDFLAGS: -static-libgcc -static-libstdc++ -Wl,-Bstatic -lstdc++ -lpthread -Wl,-Bdynamic
//#cgo windows,386 LDFLAGS: -Llib/x86 -lTolk
//#cgo windows,amd64 LDFLAGS: -Llib/x64 -lTolk
import "C"
