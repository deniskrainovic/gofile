import React from "react";
import "./App.css";
import Upload from "../Upload/Upload";
import FilePage from "../FilePage/FilePage";
import {
   BrowserRouter as Router,
   Switch,
   Route,
   Link
} from "react-router-dom";

function App() {
   return (
      <Router>
         <div className="App">
            <header className="header">
               <div className="title-container">
                  <p className="title">Go File</p>
               </div>
               <div className="verticalLine" />
               <div className="header-image" />
            </header>
            <Switch>
               <Route path="/uploads/:uploadID" component={FilePage} />
               <Route path="/">

                  <div className="MainBody">
                     <div className="Card">
                        <Upload />
                     </div>
                  </div>

               </Route>
            </Switch>
         </div>
      </Router>

   );
}



export default App;
