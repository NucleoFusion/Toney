import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'Toney',
  description: 'A Note-taking app for the terminal!',
  lang: 'en-US',
  head: [
    ['link', { rel: 'icon', href: '/favicon.ico' }],
    ['meta', { name: 'theme-color', content: '#0f172a' }]
  ],
  themeConfig: {
    outline: 'deep',
    footer: {
      message: "Released under the MIT License",
      copyright: "Created and managed by NucleoFusion",
    },
    nav: [
      { text: 'Guide', link: '/guide/' },
      { text: 'Config', link: '/config/' },
      { text: 'Install', link: '/guide/install/' },
    ],
    socialLinks: [
      { icon: 'github', link: 'https://github.com/SourcewareLab/Toney' },
      { icon: 'discord', link: 'https://discord.gg/X69MUr2DKm' }
    ],
    sidebar: {
      '/guide/': [
        { text: 'Introduction', link: '/guide/' },
        { text: 'Installation', link: '/guide/install' },
        { text: 'Contribution', link: '/guide/contribute' }
      ],
      '/config/': [
        { text: 'Introduction', link: '/config/' },
        { text: 'General', link: '/config/general' },
        {
          text: 'Styles', items: [
            { text: 'General', link: '/config/styles/' },
            { text: 'Icons', link: '/config/styles/icons' },
            { text: 'Task Styles', link: '/config/styles/taskstyles' },
            { text: 'Renderer', link: 'config/styles/renderer' }
          ]
        },
        {
          text: 'Keybinds', items: [
            { text: 'General', link: '/config/keybinds/' },
            { text: 'Global', link: '/config/keybinds/global' },
            { text: 'Home', link: '/config/keybinds/home' },
            { text: 'Daily', link: '/config/keybinds/daily' },
          ]
        }
      ]
    },
    lastUpdated: {
      text: "Last updated",
    },
    search: {
      provider: "local",
    },
  }
})
