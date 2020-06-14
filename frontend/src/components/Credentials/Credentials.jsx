import React, { useState } from "react";
import axios from "axios";
import { backendUrl } from "../../static/urls";

const Credentials = ({ customHandler, closeModalHandler }) => {
   const [password, setPassword] = useState("");
   const [expDate, setExpDate] = useState(1);
   const [downloadLink, setDownloadLink] = useState("");
   axios.defaults.withCredentials = true;

   const handleSubmit = async (e) => {
      e.preventDefault();
      axios
         .post(backendUrl + "/api/link/generate", {
            password,
            expirationDays: expDate,
         })
         .then((res) =>
            setDownloadLink(
               res.data.link.split("/")[3] + "/" + res.data.link.split("/")[4]
            )
         )
         .catch((err) => console.log(err));
      console.log("send");
      await customHandler();
   };
   return (
      <div>
         <form onSubmit={handleSubmit}>
            <label>
               Password:
               <input
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
               />
            </label>
            <label>
               Exp. Date:
               <input
                  value={expDate}
                  onChange={(e) => setExpDate(e.target.value)}
               />
            </label>
            <input type="submit" value="Create Link" />
         </form>
         <label>
            Link:{" "}
            {downloadLink !== "" ? (
               <a href={`${downloadLink}`}>Download</a>
            ) : null}
         </label>
      </div>
   );
};

export default Credentials;
