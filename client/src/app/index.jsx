import React from "react";
import classnames from "classnames";
import Sidebar from "../components/Sidebar";
import FileView from "../pages/FileView";
import * as styles from "./index.module.css";

const App = () => (
  <div
    className={classnames(styles.app, {
      [styles.themeDark]: true,
      [styles.themeLight]: false,
    })}
  >
    <Sidebar />
    <FileView />
  </div>
);

export default App;
