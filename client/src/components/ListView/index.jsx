import React from "react";
import * as styles from "./index.module.css";
import { fields } from "../../constants/listView";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faThLarge } from "@fortawesome/free-solid-svg-icons";

const ListView = () => (
  <div className={styles.listView}>
    <div className={styles.fields}>
      {fields.map((field) => (
        <div>{field}</div>
      ))}
      <FontAwesomeIcon icon={faThLarge} className={styles.displayTypeIcon} />
    </div>
  </div>
);

export default ListView;
