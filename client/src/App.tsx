import { useState } from "react";
import { Auth } from "./components/Auth";
import { Main } from "./Pages/Main";

const App = () => {
    const [statusAuth, setStatusAuth] = useState(false);
        
    return (
        <>
            { !statusAuth ? <Auth setStatus={setStatusAuth}/> : null}
            { statusAuth ? <Main/> : null }
        </>
    );
}

export default App;
