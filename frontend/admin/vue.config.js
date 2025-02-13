// const path = require('path')
// const StylelintPlugin = require('stylelint-webpack-plugin')

module.exports = {
    lintOnSave: true,
    // transpileDependencies: [
    //   'vuestic-ui',
    //   'epic-spinners',
    // ],
    // pages: {
    //   index: {
    //     // entry for the page
    //     entry: 'src/app/main.ts',
    //     // the source template
    //     template: 'public/index.html',
    //     // output as dist/index.html
    //     filename: 'index.html',
    //     // when using title option,
    //     // template title tag needs to be <title><%= htmlWebpackPlugin.options.title %></title>
    //     title: 'Vuestic Admin',
    //     // chunks to include on this page, by default includes
    //     // extracted common chunks and vendor chunks.
    //     chunks: ['chunk-vendors', 'chunk-common', 'index'],
    //   },
    // },
    configureWebpack: {
        externals: {
            "BMap": "BMap"
        },
        resolve: {
            alias: {
                // vue$: 'vue/dist/vue.esm.js',
                'vue-i18n': 'vue-i18n/dist/vue-i18n.cjs.js'
                    // '@': path.resolve('src'),
            },
        },
        module: {
            rules: [{
                test: /\.mjs$/,
                include: /node_modules/,
                type: "javascript/auto"
            }, ]
        }
        // plugins: [
        //   ...(
        //     (!lintOnSave && process.env.NODE_ENV === 'development') ? [] : [new StylelintPlugin({
        //       files: ['src/**/*.{vue,htm,html,css,sss,less,scss}'],
        //     })]
        //   ),
        // ],
    },
    css: {
        loaderOptions: {
            sass: {
                // Preload vuestic-ui variables and mixins for every component
            },
        },
    },
    pwa: {
        workboxPluginMode: 'InjectManifest',
        workboxOptions: {
            swSrc: './src/service-worker.js',
            importWorkboxFrom: 'local',
        },
    },
    devServer: {
        // open: true,    //是否自动打开浏览器
        port: 8080,
        https: false, //是否开启https
        hotOnly: false,
        proxy: {
            // 配置跨域
            '/api': {
                target: 'http://localhost:8080',
                ws: true,
                changOrigin: true, //是否开启代理
                // pathRewrite: { //  /api开头的请求会去到target下请求
                //     '^/api': '' //   替换/api 为空字符
                // }
            }
        },
    },

}