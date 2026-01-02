import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { getInfo } from "./utils/client.ts";
import { Amplify } from "aws-amplify";

const config = await getInfo();

Amplify.configure({
  Auth: {
    Cognito: {
      identityPoolId: "",
      userPoolId: config.userPoolId ?? "",
      userPoolClientId: config.userPoolWebClientId ?? "",
    },
  },
});

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App />
  </StrictMode>,
)
