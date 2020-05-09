import React, { useState } from "react";
import Dropzone from "react-dropzone";
import "./Dropzone.css";

const MyDropzone = () => {
   const [file, setFile] = useState({});
   const [fileSelected, setFileSelected] = useState(false);

   const uploadHandler = () => {
      if (fileSelected) console.log(file);
   };

   return (
      <React.Fragment>
         <Dropzone
            multiple={false}
            onDrop={(acceptedFiles) => {
               setFile(acceptedFiles);
               setFileSelected(true);
            }}
         >
            {({ getRootProps, getInputProps }) => (
               <div {...getRootProps()} className="dropzone">
                  <input {...getInputProps()} />
                  <div>Drag and drop files or upload</div>
               </div>
            )}
         </Dropzone>
         {fileSelected ? <p>File: {`${file[0].name}`}</p> : null}
         <button
            className="upload-button"
            type="button"
            disabled={!fileSelected}
            onClick={uploadHandler}
         >
            Generate Link
         </button>
      </React.Fragment>
   );
};

export default MyDropzone;
