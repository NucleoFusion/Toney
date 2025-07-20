---
layout: home
pageClass: home-page

hero:
  name: Toney 
  text: A Terminal Based Note Taking App 
  actions:
    - theme: brand 
      text: Get Started 
      link: /docs
    - theme: alt
      text: Documentation 
      link: /guide/resources

features:
  - title: Terminal-First Note Taking! 
    details:
      Take notes directly in your terminal. Designed for speed and minimal distraction, so you stay in flow while you work.
  - title: Markdown, but Better! 
    details:
      Write in plain Markdown, render beautifully. Lists, Checkboxes, all supported out of the box.
  - title: Plan Your Day! 
    details:
      Built-in daily task manager with  productivity-focused design — all inside the shell.
  - title: Single Config, Endless Possibility
    details:
      Centralized TOML config file controls everything — styling, shortcuts, layout, and more. 
---

<style>

.home-page .VPFeatures .VPFeature {
  border: 1px solid #b4befe; /* subtle bluish-gray border */
  border-radius: 12px;
  background-color: rgba(30, 32, 48, 0.4); /* semi-transparent background */
  box-shadow: 0 4px 14px rgba(125, 207, 255, 0.06); /* soft blue glow */
  transition: border-color 0.3s ease, box-shadow 0.3s ease;

}

.home-page .VPFeatures .VPFeature:hover {
  box-shadow: 0 6px 20px rgba(125, 207, 255, 0.1);
}
:root {
  /* === Hero Title Gradient (Indigo to Cyan) === */
  --vp-home-hero-name-color: transparent;
  --vp-home-hero-name-background: -webkit-linear-gradient(
    135deg,
    #b4befe,   /* Soft Purple */
    #89b4fa  /* Light Blue */
  );

  /* === Hero Image Glow Background === */
  --vp-home-hero-image-background-image: linear-gradient(
    -45deg,
    #1a1b26, /* Base background */
    #b4befe  /* Glow */
  );

  --vp-home-hero-image-filter: blur(56px);
  --overlay-gradient: color-mix(in srgb, #b4befe, transparent 80%);
}

.home-page {
  background:
    linear-gradient(225deg, var(--overlay-gradient), transparent 40%),
    radial-gradient(var(--overlay-gradient), transparent 60%) no-repeat -40vw -20vh / 120vw 180vh,
    radial-gradient(var(--overlay-gradient), transparent 70%) no-repeat 50% calc(100% + 20rem) / 60rem 30rem;

  .VPFeature code {
    background-color: #1e2030; /* Slightly lighter than main bg */
    color: #7dcfff;
    padding: 3px 6px;
    border-radius: 4px;
  }

  /* === Transparent Footer === */
  .VPFooter {
    background-color: transparent !important;
    border: none;
  }

  /* === Frosted Glass NavBar === */
  .VPNavBar:not(.top) {
    background-color: rgba(30, 32, 48, 0.6) !important;
    -webkit-backdrop-filter: blur(14px);
    backdrop-filter: blur(14px);
  }
}

/* === Responsive Hero Blur Tweaks === */
@media (min-width: 640px) {
  :root {
    --vp-home-hero-image-filter: blur(64px);
  }
}
@media (min-width: 960px) {
  :root {
    --vp-home-hero-image-filter: blur(72px);
  }
}
</style>
