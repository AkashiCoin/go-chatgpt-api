import React from 'react'
import ReactDOM from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom';
import App from './App.tsx'
import './index.css'
import {UserProvider} from "@/context/User";
import {StatusProvider} from "@/context/Status";
import Header from "@/components/Header.tsx";
import {Container} from "semantic-ui-react";
import {ToastContainer} from "react-toastify";
import Footer from "@/components/Footer.tsx";

interface WebApp {
  cdn: string;
  api: string;
  base_path: string;
  main_color: string;
}

declare global {
  interface Window {
    WEBAPP: WebApp; // 你可以替换 `any` 为 `WEBAPP` 的实际类型
  }
}


ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <UserProvider>
      <StatusProvider>
        <BrowserRouter>
          <Header/>
          <Container className={'main-content'}>
            <App />
          </Container>
          <ToastContainer />
          <Footer/>
        </BrowserRouter>
      </StatusProvider>
        </UserProvider>
  </React.StrictMode>,
)
