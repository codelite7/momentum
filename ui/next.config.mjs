/** @type {import('next').NextConfig} */
const nextConfig = {
    reactStrictMode: true,
    swcMinify: true,
    compiler: {
        relay: {
            src: "./",
            language: "typescript",
            artifactDirectory: "__generated__",
            eagerEsModules: false
        },
    },
};

export default nextConfig;
