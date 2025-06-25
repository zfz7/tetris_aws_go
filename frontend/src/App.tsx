import { useAuthenticator } from "@aws-amplify/ui-react";
import "@aws-amplify/ui-react/styles.css";
import {
  fetchUserAttributes,
  FetchUserAttributesOutput,
  signOut,
} from "aws-amplify/auth";
import { useEffect, useState } from "react";
import "./App.css";
import { LoginWindow } from "./components/Login";
import { getHello } from "./utils/client";

function App() {
  const [showLoginScreen, setShowLoginScreen] = useState(false);
  const [attributes, setAttributes] = useState<FetchUserAttributesOutput>({});
  const { authStatus, user } = useAuthenticator();

  useEffect(() => {
    const loadAttributes = async () => {
      try {
        const result: FetchUserAttributesOutput = await fetchUserAttributes();
        setAttributes(result);
      } catch (e) {
        // Unauthenticated
        setAttributes({});
      }
    };
    loadAttributes();
    if (authStatus === "authenticated") {
      setShowLoginScreen(false);
    }
  }, [authStatus, user]);

  const HomeScreen = () => {
    const [input, setInput] = useState("");
    const [output, setOutput] = useState("");
    return (
      <div className="App-header">
        <h1>Tetris template</h1>
        <h4>
          Hello {attributes.given_name} {attributes.family_name}
        </h4>
        {authStatus === "authenticated" ? (
          <button
            onClick={async () => {
              await signOut();
            }}
          >
            Sign out
          </button>
        ) : (
          <button onClick={() => setShowLoginScreen(true)}>Sign in</button>
        )}
        <div>Send to server:</div>
        <input onChange={(event) => setInput(event.target.value)}></input>
        <br />
        <button
          onClick={() =>
            getHello({ name: input }).then(
              (res) => setOutput(res.message!),
              (err) => console.log(err),
            )
          }
        >
          click me
        </button>
        <div>Server response</div>
        <div>{output}</div>
      </div>
    );
  };

  return showLoginScreen ? <LoginWindow /> : <HomeScreen />;
}

export default App;
