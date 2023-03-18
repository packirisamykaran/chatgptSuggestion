import logo from "./logo.svg";
import "./App.css";
import { useEffect, useState } from "react";
import axios from "axios";

function App() {
    const [query, setquery] = useState("");
    const [suggestions, setSuggestions] = useState([]);

    async function getSuggestions(e) {
        setquery(e.target.value);
        let result = await axios.post(
            process.env.REACT_APP_BACKEND_ENDPOINT + "getSuggestion",

            { query: e.target.value }
        );

        console.log(result.data);
        if (result.data.suggestions) {
            setSuggestions(result.data.suggestions);
        } else {
            setSuggestions([]);
        }
    }

    async function selectSuggestion(e) {
        setSuggestions([]);
        setquery(document.getElementById(e.target.id).getAttribute("value"));
    }

    return (
        <div className="App">
            <input
                type="text"
                onChange={(e) => getSuggestions(e)}
                value={query}
            />
            {suggestions.map((i, sugg) => {
                return (
                    <div
                        id={"sugg" + i}
                        value={i}
                        onClick={(e) => selectSuggestion(e)}>
                        {i}
                    </div>
                );
            })}

            <div>result</div>
        </div>
    );
}

export default App;
