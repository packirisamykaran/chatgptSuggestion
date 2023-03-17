import logo from "./logo.svg";
import "./App.css";
import { useEffect } from "react";
import axios from "axios";

function App() {
    useEffect(() => {
        let result;

        (async function () {
            result = await axios.post(
                process.env.REACT_APP_BACKEND_ENDPOINT + "getSuggestion",

                { result: "tesst" }
            );

            console.log(result);
        })();
    });

    return <div className="App">{process.env.REACT_APP_BACKEND_ENDPOINT}</div>;
}

export default App;
