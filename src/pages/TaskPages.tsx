import { useSearchParams } from "react-router-dom";
import { ArrowLeftIcon } from "lucide-react";
import { useNavigate } from "react-router-dom";

function TaskPages() {
    const [searchParams] = useSearchParams();
    const navigate = useNavigate();

    const title = searchParams.get("title");
    const description = searchParams.get("description");

    return (
        <div className="min-h-screen bg-slate-500 p-6">
            <h1 className="text-2xl font-bold text-white text-center mb-4">
                Detalhes da Tarefa
            </h1>

            <div className="bg-fuchsia-950 rounded-md p-4 shadow">
                <h2 className="text-white font-bold">
                    {title}
                </h2>

                <p className="text-white">
                    {description}
                </p>

                <button
                    onClick={() => navigate("/")}
                    className="bg-white text-slate-500 p-1 rounded-md block mx-auto mt-2"
                >
                    <ArrowLeftIcon size={18} />
                </button>
            </div>
        </div>
    );
}

export default TaskPages;