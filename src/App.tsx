import { add } from "hello-wasm";
import "./App.css";

function App() {
  return <>{add(BigInt(10), BigInt(10)).toString()}</>;
}

export default App;
