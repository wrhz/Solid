import { useEffect } from "react";

function App() {
    useEffect(() => {
        console.log("Hello React")
    }, []);

    return <h1>Hello React</h1>;
}

export default App;