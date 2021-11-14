import React from "react";
import * as styles from "./index.module.css";

const Sidebar = () => (
  <div className={styles.sidebar}>
    <div style={{ marginBottom: "15px" }}>All Files</div>
    <div>Recents</div>
  </div>
);

export default Sidebar;
