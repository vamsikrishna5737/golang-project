import React, { useEffect, useState} from "react";
import { useStateValue } from "../context/StateProvider";
import { actionType } from "../context/reducer";
import { Link } from "react-router-dom";
import Singleaddress from "./Singleaddress";

const AllProducts = () => {
  const [state, dispatch] = useStateValue();
  const [bool,setBool]=useState("false")
  const [prod,setProd] = useState({"name":"","cost":""})

  const fetchData = async () => {
    const jsonData = await fetch(process.env.REACT_APP_API + "/userproduct/"+state.user , {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });
    const products =await jsonData.json()
    console.log(products)
    if (products.message === "success"){
    dispatch({ type: actionType.ADD_Products, payload: { products: products.products}})
    }
  }

  useEffect(() => {
    fetchData();
    // eslint-disable-next-line
  }, []);


  return (
    <div className="mainpage">
      <Link to="/add">add product</Link>
        <table border={5}>
          <thead>
            <tr>
              <th className="name">ProductName</th>
              <th className="designation">Cost</th>
              <th className="company">User Email</th>
              <th className="edit">edit</th>
              <th className="delete">delete</th>
            </tr>
          </thead>
          <tbody>
            {state.products.map((obj,idx) => (
              <Singleaddress key={idx} obj={obj} fetchData={fetchData}/>
            ))}
          </tbody>
        </table>
    </div>
  );
};

export default AllProducts;
