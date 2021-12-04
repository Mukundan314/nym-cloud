import React, { useState } from "react";
import * as styles from "./index.module.css";
import SearchBar from "../../components/SearchBar";
import Profile from "../../components/Profile";
import Upload from "../../components/Upload";
import ListView from "../../components/ListView";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faQuestionCircle,
  faBell,
  faCog,
} from "@fortawesome/free-solid-svg-icons";

const FileView = () => {
  const [search, setSearch] = useState("");

  return (
    <div className={styles.file}>
      <div className={styles.toolBar}>
        <SearchBar search={search} setSearch={setSearch} />
        <div className={styles.options}>
          <div>
            <FontAwesomeIcon
              icon={faQuestionCircle}
              className={styles.optionsIcon}
            />
            <FontAwesomeIcon icon={faCog} className={styles.optionsIcon} />
            <FontAwesomeIcon icon={faBell} className={styles.optionsIcon} />
          </div>
          <Profile />
        </div>
      </div>
      <div className={styles.sectionHeader}>
        <div className={styles.sectionName}>All Files</div>
        <Upload />
      </div>
      <ListView search={search} />
    </div>
  );
};

export default FileView;
