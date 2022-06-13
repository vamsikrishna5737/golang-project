import "./signin.css";
import React, { useState, useRef } from "react";
import { Link, useNavigate } from "react-router-dom";
import { useStateValue } from "../context/StateProvider";
import { actionType } from "../context/reducer";
import Alert from "../components/Alert";
import dot from "../utils/dot.svg";
import bigCircleL from "../utils/bigCircleL.svg";
import bigCircleR from "../utils/bigCircleR.svg";
import eye from "../utils/eye.svg";
import { motion } from "framer-motion";

const Signin = () => {
  const navigate = useNavigate();
  const [msg, setMsg] = useState(null);
  const passRef = useRef();
  const [state, dispatch] = useStateValue();

  const handleSubmit = async (e) => {
    e.preventDefault();
    const data = {
      email: e.target.elements.log_email.value,
      password: e.target.elements.log_password.value,
    };

    const proRes = await fetch(process.env.REACT_APP_API + "/user", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    });
    const response = await proRes.json();
    if (response.message === "success") {
      console.log(response.data.email)
      console.log(response.token)
      dispatch({ type: actionType.ADD_USER, payload: { user: response.data.email}})
      dispatch({ type: actionType.ADD_Token, payload: { token: response.token}})
      navigate("/product");
    } else {
      setMsg(response.message);
      setTimeout(() => {
        setMsg(null);
      }, 2500);
    }
  };

  function showPassword(e) {
    var x = passRef.current;
    if (x.type === "password") {
      x.type = "text";
    } else {
      x.type = "password";
    }
  }

  return (
    <section className="loginContainer">
      <img src={bigCircleL} alt="bigCircle" className="bigCircle left" />
      <div className="mainLogIn">
        <img src={dot} alt="dotLeft" className="dotLeft" />

        {msg ? <Alert msg={msg} /> : ""}

        <form action="" className="signForm" onSubmit={handleSubmit}>
          <div className="logo">Logo</div>

          <div className="detail">
            Enter your credentials to access your account
          </div>
          <input
            type="email"
            placeholder="User Id"
            id="log_email"
            className="inputauth"
            required
          />
          <input
            type="password"
            placeholder="Password"
            id="log_password"
            className="inputauth"
            required
            ref={passRef}
          />
          <motion.img
            whileTap={{ scale: 0.95 }}
            src={eye}
            alt="eye"
            onClick={showPassword}
            className="signineye"
          />
          <motion.button
            whileTap={{ scale: 0.95 }}
            type="submit"
            className="btn"
          >
            Sign In
          </motion.button>
          <Link to="/signup" className="linkLogin">
            Sign up
          </Link>
        </form>

        <img src={dot} alt="dotRight" className="dotRight" />
      </div>
      <img src={bigCircleR} alt="bigCircle" className="bigCircle right" />
    </section>
  );
};

export default Signin;
