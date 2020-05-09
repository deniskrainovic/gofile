import React from "react";
import "./App.css";
import MyDropzone from "../Dropzone/Dropzone";

function App() {
   return (
      <div className="App">
         <header className="header">
            <div>
               <p>Go File</p>
            </div>
            <div className="verticalLine" />
         </header>
         <MyDropzone />
      </div>
   );
}

export default App;
