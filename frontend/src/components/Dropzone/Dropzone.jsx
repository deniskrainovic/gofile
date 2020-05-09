import React, { useState } from "react";
import Dropzone from "react-dropzone";
import "./Dropzone.css";

const MyDropzone = () => {
   const [file, setFile] = useState({});
   const [fileSelected, setFileSelected] = useState(false);

   const uploadHandler = () => {
      console.log(file);
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
                  <div>
                     Drag 'n' drop some files here, or click to select files
                  </div>
               </div>
            )}
         </Dropzone>
         {fileSelected ? <p>File: {`${file[0].name}`}</p> : null}
         <button type="button" onClick={uploadHandler}>
            Upload
         </button>
      </React.Fragment>
   );
};

export default MyDropzone;
