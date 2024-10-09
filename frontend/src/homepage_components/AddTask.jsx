import addIcon from '../homepage_components../src./public/add_icon.png'
function AddTask() {
    return(
    <div className="add-task-div">
        <button className="add-task-button">
            <img src={addIcon} alt="add task button" height="25px"></img>
        </button>
    </div>
    );
}

export default AddTask;