import React, { Component } from "react";
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link,
    useParams
  } from "react-router-dom";
import "./FilePage.css";

class Dropzone extends Component {
    constructor(props) {
        super(props);
        this.state = { mode: undefined } ;
      }

    componentWillMount(){
        const { uploadID } = this.props.match.params
        let xhr = new XMLHttpRequest();
        xhr.open('POST', 'http://localhost:8080/uploads/' + uploadID);

        xhr.setRequestHeader('Content-type', 'application/json');
    }
  
    render() {
      return (
      <h1>Test Page Uplaod </h1>
      );
    }
  }
  
  export default Dropzone;