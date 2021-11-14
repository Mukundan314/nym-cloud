import React, { useState } from "react";
import classnames from "classnames";
import Sidebar from "../components/Sidebar";
import FileView from "../pages/FileView";
import * as styles from "./index.module.css";

const App = () => {
  const [tab, setTab] = useState("all-files");

  return (
    <div
      className={classnames(styles.app, {
        [styles.themeDark]: true,
        [styles.themeLight]: false,
      })}
    >
      <Sidebar tab={tab} setTab={setTab} />
      <FileView />
    </div>
  );
};
export default App;
