import React from "react";
import { useDropzone } from "react-dropzone";

const MyDropzone = ({ onFilesAdded }) => {
   const onDrop = (acceptedFiles) => {
      onFilesAdded(acceptedFiles);
   };

   const { getRootProps, getInputProps, isDragActive } = useDropzone({
      onDrop,
   });

   return (
      <div className="Dropzone" {...getRootProps()}>
         <img
            alt="upload"
            className="Icon"
            src="baseline-cloud_upload-24px.svg"
         />
         <input {...getInputProps()} multiple type="file" />
         {isDragActive ? (
            <p>Drop the files here ...</p>
         ) : (
            <p>Drag 'n' drop some files here, or click to select files</p>
         )}
      </div>
   );
};

export default MyDropzone;
