// const getUser = () => {
//   return localStorage.getItem("user")
//     ? JSON.parse(localStorage.getItem("user"))
//     : { email: "NA", token: "NA" };
// };

const initialState = {
  user: "",
  token:"",
  products:[],
  id:null,
};

export default initialState;
