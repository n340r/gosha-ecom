
// eslint-disable-next-line @typescript-eslint/no-var-requires
const { composePlugins, withNx } = require('@nx/next');

const isProd = process.env.NODE_ENV === 'production';

// CSP
const csp = [
  "default-src 'self'",
  "base-uri 'self'",
  "object-src 'none'",
  "frame-ancestors 'none'",
  `script-src 'self' https: 'unsafe-inline'${isProd ? '' : " 'unsafe-eval'"}`,
  "style-src 'self' 'unsafe-inline' https:",
  "img-src 'self' data: blob: https: s3-s1.retailcrm.tech",
  `connect-src 'self' ${process.env.RETAILCRM_BASE_URL || ''} ${process.env.RETAILCRM_TEST_BASE_URL || ''} https://api.yookassa.ru https://api.cdek.ru https://api.telegram.org https://cdn.jsdelivr.net wss:`,
  "frame-src 'none'",
  "font-src 'self' data: https:",
  "worker-src 'self' blob:",
  "form-action 'self'",
  "upgrade-insecure-requests"
].join('; ');

/** @type {import('@nx/next/plugins/with-nx').WithNxOptions & import('next').NextConfig} */
const nextConfig = {
  nx: {},

  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 's3-s1.retailcrm.tech',
        pathname: '/ru-central1/retailcrm/**'
      }
    ]
  },

  webpack(config) {
    config.module.rules.push({
      test: /\.svg$/,
      use: ['@svgr/webpack']
    });
    return config;
  },

  async headers() {
    const common = [
      { key: 'Content-Security-Policy', value: csp },
      { key: 'Referrer-Policy', value: 'strict-origin-when-cross-origin' },
      { key: 'X-Content-Type-Options', value: 'nosniff' },
      { key: 'X-Frame-Options', value: 'DENY' },
      { key: 'Permissions-Policy', value: 'camera=(), microphone=()' },
      { key: 'Cross-Origin-Opener-Policy', value: 'same-origin' },
      { key: 'Cross-Origin-Resource-Policy', value: 'same-origin' },
      { key: 'X-DNS-Prefetch-Control', value: 'off' }
    ];

    if (isProd) {
      common.push({
        key: 'Strict-Transport-Security',
        value: 'max-age=31536000; includeSubDomains; preload'
      });
    }

    return [{ source: '/:path*', headers: common }];
  }
};

module.exports = composePlugins(withNx)(nextConfig);
