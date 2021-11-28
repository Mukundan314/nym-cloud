import React from "react";
import * as styles from "./index.module.css";
import ListItem from "../ListItem";
import { fields } from "../../constants/listView";
import { files } from "../../constants/files";
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
    <div>
      {files.map((file, idx) => (
        <div className={styles.listItem}>
          <ListItem
            name={file.name}
            date={file.date}
            size={file.size}
            type={file.type}
            idx={idx}
          />
        </div>
      ))}
    </div>
  </div>
);

export default ListView;
