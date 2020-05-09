import React from "react";
import "./App.css";
import MyDropzone from "../Dropzone/Dropzone";

function App() {
   return (
      <div className="App">
         <header className="header">
            <div className="title-container">
               <p className="title">Go File</p>
            </div>
            <div className="verticalLine" />
            <div className="header-image" />
         </header>
         <MyDropzone />
      </div>
   );
}

export default App;
