/** @type {import('next').NextConfig} */
const nextConfig = {
  experimental: {
    instrumentationHook: true,
  },
  reactStrictMode: false,
  swcMinify: true,
  compiler: {
    relay: {
      src: "./",
      language: "typescript",
      artifactDirectory: "__generated__",
      eagerEsModules: false,
    },
  },
};

module.exports = nextConfig;
