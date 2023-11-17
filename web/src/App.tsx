import {Suspense} from "react";
import Home from "@/components/Home.tsx";
import 'semantic-ui-css/semantic.min.css';
import 'react-toastify/dist/ReactToastify.css';
import './App.css'
import {Routes, Route} from "react-router-dom";
import Loading from "@/components/Loading.tsx";

function App() {
  return (
      <Routes>
          <Route
              path='/'
              element={
                  <Suspense fallback={<Loading></Loading>}>
                      <Home />
                  </Suspense>
              }
          />
      </Routes>
  )
}

export default App
