import type { NextConfig } from "next";

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

export default nextConfig;