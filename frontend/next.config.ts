import type { NextConfig } from "next";
import { withSentryConfig } from "@sentry/nextjs";

const nextConfig: NextConfig = {
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'sspb.digitaistudios.com',
        pathname: '/api/files/**',
      },
    ],
  },
};

// Wrap config with Sentry
export default withSentryConfig(nextConfig, {
  // Opcions de Sentry
  silent: true, // Redueix log en build
  org: "digitaistudios", // Canvia al teu org de Sentry
  project: "crims-frontend", // Canvia al teu projecte de Sentry

  // Opcions de sourcemaps
  widenClientFileUpload: true,
  tunnelRoute: "/monitoring",

  // Opcions de webpack
  hideSourceMaps: true,
});
