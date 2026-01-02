import "@testing-library/jest-dom";
import { Amplify } from "aws-amplify";
import { vi } from "vitest";

// Mock Amplify configuration for tests
Amplify.configure({
  Auth: {
    Cognito: {
      identityPoolId: "test-identity-pool-id",
      userPoolId: "test-user-pool-id",
      userPoolClientId: "test-user-pool-client-id",
    },
  },
});

// Mock AWS Amplify auth functions
vi.mock("aws-amplify/auth", () => ({
  fetchUserAttributes: vi.fn().mockResolvedValue({
    given_name: "John",
    family_name: "Doe",
    email: "john.doe@example.com",
    sub: "mock-user-id-123",
  }),
  fetchAuthSession: vi.fn().mockResolvedValue({
    tokens: {
      idToken: {
        toString: () => "mock-token",
        payload: {
          exp: Date.now() / 1000 + 3600, // 1 hour from now
        },
      },
    },
  }),
  signOut: vi.fn().mockResolvedValue(undefined),
}));

// Mock the Authenticator hook
vi.mock("@aws-amplify/ui-react", async () => {
  const actual = await vi.importActual("@aws-amplify/ui-react");
  return {
    ...actual,
    useAuthenticator: vi.fn(() => ({
      authStatus: "authenticated",
      user: {
        userId: "mock-user-id-123",
        username: "john.doe@example.com",
        attributes: {
          given_name: "John",
          family_name: "Doe",
          email: "john.doe@example.com",
          sub: "mock-user-id-123",
        },
      },
    })),
  };
});
