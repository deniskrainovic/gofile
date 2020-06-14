import React, { useState } from "react";
import Modal from "react-modal";
import Credentials from "../Credentials/Credentials";

const customStyles = {
   content: {
      top: "50%",
      left: "50%",
      right: "auto",
      bottom: "auto",
      marginRight: "-50%",
      transform: "translate(-50%, -50%)",
   },
};

// Make sure to bind modal to your appElement (http://reactcommunity.org/react-modal/accessibility/)
Modal.setAppElement("#root");

const MyModal = (props) => {
   const [modalIsOpen, setIsOpen] = useState(false);
   const openModal = () => {
      setIsOpen(true);
   };

   const closeModal = () => {
      setIsOpen(false);
   };

   return (
      <div>
         <button onClick={openModal}>Generate Link</button>
         <Modal
            isOpen={modalIsOpen}
            onRequestClose={closeModal}
            style={customStyles}
            contentLabel="Credentials"
         >
            <Credentials
               customHandler={props.customHandler}
               closeModalHandler={closeModal}
            />
         </Modal>
      </div>
   );
};

export default MyModal;
