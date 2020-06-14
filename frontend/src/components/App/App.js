import React from "react";
import "./App.css";
import Upload from "../Upload/Upload";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import UploadedFiles from "../UploadedFiles/UploadedFiles.jsx";

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
               <Route exact path="/">
                  <div className="MainBody">
                     <div className="Card">
                        <Upload />
                     </div>
                  </div>
               </Route>
               <Route path="/uploads/:id" component={UploadedFiles}></Route>
            </Switch>
         </div>
      </Router>
   );
}

export default App;
