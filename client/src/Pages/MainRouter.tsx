import {BrowserRouter, Routes, Route} from "react-router-dom";
import { Auth } from "../components/Auth";
import { Main } from "./Main";
import history from "../Services/history";


export const MainRouter = () => {
    // TODO Сделать защищеные роуты


    console.log("history",history);
    
      return (
          <BrowserRouter>
              <Routes>
                  {/* <Route path="/" element={<Auth/>}/>
                  <Route path="/#/main" element={<Main/>}/> */}
              </Routes>
          </BrowserRouter>
      )
}