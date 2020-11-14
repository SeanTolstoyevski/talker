Kurulum, Fonksiyonlar, Kullanım ve derleme sonrası aşamalar
-----------------------------------------------------------

Kurulum
~~~~~~~

Projenin **GCCGo** bağımlılığı olduğu için Go’nun **gcc**’ye
ulaşabiliyor olması gerekir.

::

   go get -tags bit64 github.com/SeanTolstoyevski/talker

Herhangi bir çıktı görmediğinizde modül kurulmuş demektir.

Fonksiyonlar
~~~~~~~~~~~~

-  ``talker.Load()``

Tolk’u yükler ve yapılandırır. Diğer fonksiyonları çalıştırmadan önce bu
fonksiyonu çalıştırdığınıza emin olmalısınız.

-  Dönüş: hiçbir şey.

C açıklaması: Initializes Tolk by loading and initializing the screen
reader drivers and setting the current screen reader driver, provided at
least one of the supported screen readers is active. Also initializes
COM if it has not already been initialized on the calling thread.
Calling this function more than once will only initialize COM. You
should call this function before using the functions below. Use
Tolk_IsLoaded to determine if Tolk has been initialized.

-  ``talker.Unload()``

Tolk’un kaynaklarını serbest bırakır. Uygulamalardan çıkma gibi
durumlarda çalıştırılması gerekir.

-  Dönüş: hiçbir şey

C açıklaması: Finalizes Tolk by finalizing and unloading the screen
reader drivers and clearing the current screen reader driver, provided
one was set. Also uninitializes COM on the calling thread. Calling this
function more than once will only uninitialize COM. You should not use
the functions below if this function has been called.

-  ``talker.IsLoaded() bool``

Tolk’un yapılandırılma durumunu kontrol eder. Eğer daha önce Tolk
yüklenmişse True, aksi türlü false döndürür

-  Dönüş: bool

C açıklaması: Tests if Tolk has been initialized.

-  ``talker.TrySAPI(yesno bool)``

SAPI’nın dahil edilip edilmeyeceiğini belirler. En iyi durumda Load’den
önce kullanmalısınız.

-  Dönüş: hiçbir şey

C açıklaması: Sets if Microsoft Speech API (SAPI) should be used in the
screen reader auto-detection process. The default is not to include
SAPI. The SAPI driver will use the system default synthesizer, voice and
soundcard. This function triggers the screen reader detection process if
needed. For best performance, you should call this function before
calling Tolk_Load.

-  ``talker.PreferSAPI(yesno bool)``

SAPI driver’inin ilk öğe veya son öğe olma durumunu ayarlar. Varsayılan
durumda SAPI son öğedir. En iyi durumda Load’den önce kullanmalısınız.

-  Dönüş: hiçbir şey

C açıklaması: If auto-detection for SAPI has been turned on through
Tolk_TrySAPI, sets if SAPI should be placed first (true) or last (false)
in the screen reader detection list. Putting it last is the default and
is good for using SAPI as a fallback option. Putting it first is good
for ensuring SAPI is used even when a screen reader is running, but keep
in mind screen readers will still be tried if SAPI is unavailable. This
function triggers the screen reader detection process if needed. For
best performance, you should call this function before calling
Tolk_Load.

-  ``talker.DetectScreenReader() string``

Şu an aktif olan ekran okuyucusunun adını döndürür (Örn. NVDA, JAWS…)

-  Dönüş: string

C açıklaması: Returns the common name for the currently active screen
reader driver, if one is set. If none is set, tries to detect the
currently active screen reader before looking up the name. If no screen
reader is active, NULL is returned. Note that the drivers hard-code the
common name, it is not requested from the screen reader itself. You
should call Tolk_Load once before using this function.

-  ``talker.Output(text string, interrupt bool) bool``

Metni aktif olan ekran okuyucusuna gönderir. Eğer “interrupt” true ise
mevcut konuşma duraklatılır. Çoğu zaman “interrupt”un true olmasına
gerek yoktur. Kullanıcılar mevcut konuşmanın kesilmesinden hoşlanmazlar.

-  Dönüş: bool

C açıklaması: Outputs text through the current screen reader driver, if
one is set. If none is set or if it encountered an error, tries to
detect the currently active screen reader before outputting the text.
This is the preferred function to use for sending text to a screen
reader, because it uses all of the supported output methods (speech
and/or braille depending on the current screen reader driver). You
should call Tolk_Load once before using this function. This function is
asynchronous.

