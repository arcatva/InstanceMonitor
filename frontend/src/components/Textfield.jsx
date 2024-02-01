import React from "react";

import styles from "./Textfield.module.css"

const Textfield = ({name,type,placeholder}) =>{

    return(
        <div className={styles.textfield}>
            <div className="text-xs font-medium text-gray-600"> {name} </div>
            <input
                type={type}
                id={name}
                placeholder={placeholder}
                className="mt-1 w-full border-none p-0 focus:border-transparent focus:outline-none focus:ring-0 sm:text-sm"
            />
        </div>
    )
}

export default Textfield