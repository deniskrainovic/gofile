import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";
import { backendUrl } from "../../static/urls";

const UploadedFiles = () => {
   let { id } = useParams();
   const [passwordNeeded, setPasswordNeeded] = useState(false);
   const [enteredPassword, setEnteredPassword] = useState("");
   const [passwordConfirmed, setPasswordConfirmed] = useState(false);

   const [files, setFiles] = useState([]);

   const fetchFilesHandler = (event) => {
      if (event !== undefined) event.preventDefault();
      axios
         .post(`${backendUrl}/api/uploads/${id}`, {
            password: enteredPassword,
         })
         .then((res) => {
            if (res.status === 200) {
               setFiles(res.data.files);
               setPasswordConfirmed(true);
            }
         })
         .catch((err) => console.log(err));
   };

   const downloadFilesHandler = (event) => {
      event.preventDefault();

      if (passwordNeeded) {
         axios({
            url: `${backendUrl}/api/uploads/${id}/download`,
            method: "POST",
            responseType: "blob",
            data: {
               password: enteredPassword,
            },
         })
            .then((response) => {
               const url = window.URL.createObjectURL(
                  new Blob([response.data])
               );
               const link = document.createElement("a");
               link.href = url;
               link.setAttribute("download", "goFile.zip");
               document.body.appendChild(link);
               link.click();
            })
            .catch((err) => console.log(err));
      } else {
         axios({
            url: `${backendUrl}/api/uploads/${id}/download`,
            method: "POST",
            responseType: "blob",
         })
            .then((response) => {
               const url = window.URL.createObjectURL(
                  new Blob([response.data])
               );
               const link = document.createElement("a");
               link.href = url;
               link.setAttribute("download", "goFile.zip");
               document.body.appendChild(link);
               link.click();
            })
            .catch((err) => console.log(err));
      }
   };

   let renderContent = (
      <form onSubmit={(password) => fetchFilesHandler(password)}>
         <label>Password: </label>
         <input
            type="password"
            name="password"
            required
            value={enteredPassword}
            onChange={(e) => setEnteredPassword(e.target.value)}
         />
         <input type="submit" value="Show Data"></input>
      </form>
   );
   if (passwordConfirmed) {
      renderContent = (
         <form onSubmit={(password) => downloadFilesHandler(password)}>
            <input type="submit" value="Download"></input>
         </form>
      );
   }

   useEffect(() => {
      axios
         .get(`${backendUrl}/api/uploads/${id}/checkpassword`)
         .then((res) => {
            setPasswordNeeded(res.data.isPasswordNeeded);
            return res;
         })
         .then((res) => {
            if (!res.data.isPasswordNeeded) {
               fetchFilesHandler();
            }
            return true;
         })
         .catch((err) => console.log(err));

      console.log("didmount");
   }, []);

   return (
      <div>
         <ul>
            {files.map((file, index) => (
               <li key={index}>
                  <label>{file.Originalname}</label>
                  {/* TODO: download only this file */}
               </li>
            ))}
         </ul>
         {passwordNeeded ? (
            renderContent
         ) : (
            <form onSubmit={(password) => downloadFilesHandler(password)}>
               <input type="submit" value="Download"></input>
            </form>
         )}
      </div>
   );
};

export default UploadedFiles;
