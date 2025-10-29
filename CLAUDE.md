# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

<!-- nx configuration start-->
<!-- Leave the start & end comments to automatically receive updates. -->

# General Guidelines for working with Nx

- When running tasks (for example build, lint, test, e2e, etc.), always prefer running the task through `nx` (i.e. `nx run`, `nx run-many`, `nx affected`) instead of using the underlying tooling directly
- You have access to the Nx MCP server and its tools, use them to help the user
- When answering questions about the repository, use the `nx_workspace` tool first to gain an understanding of the workspace architecture where applicable.
- When working in individual projects, use the `nx_project_details` mcp tool to analyze and understand the specific project structure and dependencies
- For questions around nx configuration, best practices or if you're unsure, use the `nx_docs` tool to get relevant, up-to-date docs. Always use this instead of assuming things about nx configuration
- If the user needs help with an Nx configuration or project graph error, use the `nx_workspace` tool to get any errors

<!-- nx configuration end-->

## Project Overview

This is an e-commerce application for SERVISEX™ built with Next.js 14 and managed as an Nx monorepo. The application handles product browsing, shopping cart management, checkout with multiple payment methods (YooKassa), and delivery integration (CDEK). Currently the site is in a suspended state displaying a maintenance page.

## Architecture

### Workspace Structure

The monorepo contains two main projects:
- **front**: Next.js application (App Router) - main e-commerce frontend
- **front-e2e**: Playwright end-to-end tests for the frontend

### Technology Stack

**Frontend Framework:**
- Next.js 14.2.6 with App Router
- React 18 with TypeScript
- Tailwind CSS for styling with custom animations

**State Management:**
- Zustand for client-side state (cart management)
- TanStack Query (React Query) for server state management
- Redux Toolkit with Redux Persist (legacy, may be phased out)

**UI Components:**
- Radix UI primitives (Dialog, Select, Tabs, Radio, Label, Separator)
- Custom UI components built with class-variance-authority
- Lucide React for icons
- Framer Motion for animations
- Swiper and Embla Carousel for image galleries

**Forms & Validation:**
- React Hook Form with Zod validation
- Custom checkout form schemas

**Third-Party Integrations:**
- **RetailCRM**: Backend system for product catalog and order management
- **YooKassa**: Payment gateway (production and test modes)
- **CDEK**: Delivery service integration
- **Telegram**: Order notifications

### Key Architectural Patterns

**API Routes Structure:**
All API routes are in `front/src/app/api/`:
- Product data: `getProducts`, `getProductsByIds`
- Order management: `createOrder`, `getOrdersByIds`, `checkOrderStatus`
- Payment: `createPayment`, `createTestPayment`, `yooKassa`, `yooKassaTest`
- Delivery: `cdek`

**Configuration:**
- Environment-aware configuration in `front/src/lib/server/config.ts`
- Supports production/development modes with separate RetailCRM and YooKassa credentials
- Server-only configuration pattern using "server-only" package

**Security:**
- Comprehensive CSP headers configured in `next.config.js`
- Security headers include HSTS, X-Frame-Options, CORS policies
- Different CSP rules for production vs development

**Component Organization:**
- UI components follow atomic design in `front/src/components/ui/`
- Business components in `front/src/components/`
- Layout components in `front/src/layouts/`
- Centralized exports via index files

**Type System:**
- Comprehensive TypeScript types in `front/src/types/`
- Separated by domain: product, cart, order, checkout, payment (yookassa), delivery (cdek), telegram

**Custom Hooks:**
- `useCartStore`: Zustand-based cart management
- `useConfirmationDialog`: Dialog state management
- `useIntro`: Intro animation/state
- `useOrderStatus`: Order status polling
- `useStore`: Hydration-safe store hook

## Development Commands

### Running the Application

```bash
# Start development server (runs on port 4200)
pnpm dev
# or
pnpm nx run front:dev

# Build for production
pnpm build
# or
pnpm nx run front:build

# Start production server (runs on port 3000)
pnpm start
# or
pnpm nx run front:start
```

### Code Quality

```bash
# Lint the frontend
pnpm lint
# or
pnpm nx run front:lint

# Type checking
pnpm typecheck
# or
pnpm nx run front:typecheck
# or
tsc -p tsconfig.json

# Type check all projects
pnpm nx run-many --target=typecheck
```

### Testing

```bash
# Run e2e tests
pnpm nx run front-e2e:e2e

# Run e2e tests in CI mode
pnpm nx run front-e2e:e2e-ci
```

### Package Management

This project uses **pnpm** as the package manager. The workspace is configured with `pnpm-workspace.yaml` and only pnpm is allowed (enforced via preinstall hook).

```bash
# Install dependencies
pnpm install

# Add a dependency to the front project
pnpm --filter front add <package-name>
```

## Important Notes

### Environment Variables

The application requires several environment variables (stored in `front/.env.local`):
- `RETAILCRM_ENV`: "production" or "development"
- `RETAILCRM_API_KEY`, `RETAILCRM_TEST_API_KEY`: RetailCRM API keys
- `RETAILCRM_PROD_BASE_URL`, `RETAILCRM_TEST_BASE_URL`: RetailCRM endpoints
- `YOOKASSA_SHOP_ID`, `YOOKASSA_KEY`: Production payment credentials
- `YOOKASSA_TEST_SHOP_ID`, `YOOKASSA_TEST_KEY`: Test payment credentials

### Current State

The application is currently displaying a maintenance/suspension page (see `front/src/app/layout.tsx:24`). The actual application is commented out but fully implemented. To enable the full application, uncomment line 23 and remove the `<StayTunedPlaceholder />` component.

### SVG Handling

SVGs are imported as React components using @svgr/webpack (configured in `next.config.js`). Type definitions are in `front/src/types/svg.d.ts`.

### Nx Plugins Configuration

The workspace uses these Nx plugins (see `nx.json`):
- `@nx/js/typescript`: TypeScript compilation with targets: typecheck, build, build-deps, watch-deps
- `@nx/next/plugin`: Next.js tasks with targets: start, build, dev, serve-static
- `@nx/playwright/plugin`: E2E testing with target: e2e
- `@nx/eslint/plugin`: Linting with target: lint
