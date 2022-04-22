run: nextjs/out/index.html \
     nuxtjs/dist/index.html \
     vite-vue/dist/index.html \
     vite-react/dist/index.html
	go run . -p :8000

distclean:
	rm -rf nuxtjs/dist
	rm -rf nextjs/out
	rm -rf vite-vue/dist
	rm -rf vite-react/dist

nextjs/out/index.html:
	cd nextjs; yarn && yarn build && yarn next export

nuxtjs/dist/index.html:
	cd nuxtjs; yarn && yarn generate

vite-vue/dist/index.html:
	cd vite-vue; yarn && yarn build

vite-react/dist/index.html:
	cd vite-react; yarn && yarn build

.PHONY: nextjs nuxtjs vite-react vite-vue
nextjs: nextjs/out/index.html
	(cd nextjs/out/ && python3 -m http.server 8000)

nuxtjs: nuxtjs/dist/index.html
	(cd nuxtjs/dist/ && python3 -m http.server 8000)

vite-react: vite-react/dist/index.html
	(cd vite-react/dist/ && python3 -m http.server 8000)

vite-vue: vite-vue/dist/index.html
	(cd vite-vue/dist/ && python3 -m http.server 8000)


