run: nextjs/out/index.html \
     nuxtjs/dist/index.html \
     vite-vue/dist/index.html \
     vite-react/dist/index.html
	go run . -p :8000

nextjs/out/index.html:
	cd nextjs; yarn && yarn build

nuxtjs/dist/index.html:
	cd nuxtjs; yarn && yarn generate

vite-vue/dist/index.html:
	cd vite-vue; yarn && yarn build

vite-react/dist/index.html:
	cd vite-react; yarn && yarn build
