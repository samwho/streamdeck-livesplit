GO = go
GOFLAGS =
NAME = dev.samwho.streamdeck-livesplit
INSTALLDIR = "$(APPDATA)\Elgato\StreamDeck\Plugins\$(NAME).sdPlugin"
BUILDDIR = build
RELEASEDIR = release
SDPLUGINDIR = "$(BUILDDIR)\$(NAME).sdPlugin"

.PHONY: test install build logs

build:
	$(GO) build $(GOFLAGS)

test:
	$(GO) run $(GOFLAGS) main.go -port 12345 -pluginUUID 213 -registerEvent test -info "{\"application\":{\"language\":\"en\",\"platform\":\"mac\",\"version\":\"4.1.0\"},\"plugin\":{\"version\":\"1.1\"},\"devicePixelRatio\":2,\"devices\":[{\"id\":\"55F16B35884A859CCE4FFA1FC8D3DE5B\",\"name\":\"Device Name\",\"size\":{\"columns\":5,\"rows\":3},\"type\":0},{\"id\":\"B8F04425B95855CF417199BCB97CD2BB\",\"name\":\"Another Device\",\"size\":{\"columns\":3,\"rows\":2},\"type\":1}]}"

sdplugin: build
	rm -rf $(SDPLUGINDIR)
	mkdir -p $(SDPLUGINDIR)
	cp *.json $(SDPLUGINDIR)
	cp *.exe $(SDPLUGINDIR)
	cp -r images $(SDPLUGINDIR)

install: sdplugin
	rm -rf $(INSTALLDIR)
	mv $(SDPLUGINDIR) $(INSTALLDIR)

uninstall:
	rm -rf $(INSTALLDIR)

logs:
	tail -f "$(TMP)"/streamdeck-livesplit.log*

release: sdplugin
	rm -rf $(RELEASEDIR)
	mkdir $(RELEASEDIR)
	DistributionTool.exe -b -i $(SDPLUGINDIR) -o $(RELEASEDIR)