import React from "react";
import { useStateValue } from "../context/StateProvider";
import user from "../utils/user.svg";
import { motion } from "framer-motion";

const Header = () => {
  const [state, dispatch] = useStateValue();
  

  return (
    <header>
      <ul className="header-ul">
        <motion.li whileTap={{ scale: 0.6 }} className="mainHead">
          Total Contact
        </motion.li>
        <li>
          <input
            type="search"
            placeholder="Search by product name"
            className="searchBox"
            // onChange={handleChange}
          />
        </li>
        <motion.li whileTap={{ scale: 0.6 }} className="userField">
          <img src={user} alt="user" />
          <div>
            <p>
              {state.user.split("@")[0]}
            </p>
            <p className="userType">Normal User</p>
          </div>
        </motion.li>
      </ul>
    </header>
  );
};

export default Header;
