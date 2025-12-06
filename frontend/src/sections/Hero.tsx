import { Calendar, MapPin } from "lucide-react";

const Hero = () => {
  return (
    <div className="relative pt-16 pb-32 flex content-center items-center justify-center min-h-[85vh]">
      {/* Background Image */}
      <div
        className="absolute top-0 w-full h-full bg-center bg-cover z-0"
        style={{
          backgroundImage:
            "url('https://images.unsplash.com/photo-1492684223066-81342ee5ff30?auto=format&fit=crop&q=80&w=1920')",
        }}
      >
        <span className="w-full h-full absolute opacity-60 bg-black"></span>
        <span className="w-full h-full absolute bg-gradient-to-t from-gray-900 via-transparent to-transparent"></span>
      </div>

      <div className="container relative mx-auto z-10 text-center px-4">
        <div className="items-center flex flex-wrap">
          <div className="w-full lg:w-8/12 px-4 ml-auto mr-auto text-center">
            <h1 className="text-white font-extrabold text-5xl md:text-6xl leading-tight mb-6">
              Vive Momentos <br />
              <span className="text-indigo-400">Inolvidables</span>
            </h1>
            <p className="mt-4 text-lg text-gray-200 mb-12 max-w-2xl mx-auto">
              Descubre y reserva los mejores conciertos, conferencias y eventos exclusivos en tu ciudad.
            </p>

            {/* Featured Event Badge */}
            <div className="inline-flex flex-col sm:flex-row items-center bg-white/10 backdrop-blur-md rounded-2xl p-4 text-left border border-white/20 shadow-2xl max-w-lg mx-auto sm:max-w-none w-auto">
              <img
                src="https://images.unsplash.com/photo-1501281668745-f7f57925c3b4?auto=format&fit=crop&q=80&w=200"
                className="w-24 h-24 rounded-xl object-cover mb-4 sm:mb-0 sm:mr-6 border-2 border-indigo-500/50"
                alt="Evento destacado"
              />
              <div className="text-white">
                <span className="inline-block px-2 py-1 text-xs font-semibold bg-indigo-600 rounded-full mb-2">
                  Destacado del Mes
                </span>
                <h3 className="text-xl font-bold mb-1">Tomorrowland 2025 - Cierre</h3>
                <div className="flex items-center text-sm text-gray-300 space-x-4">
                  <span className="flex items-center"><Calendar className="w-4 h-4 mr-1"/> 25 Jul</span>
                  <span className="flex items-center"><MapPin className="w-4 h-4 mr-1"/> Boom, BE</span>
                </div>
              </div>
              <button className="mt-4 sm:mt-0 sm:ml-6 bg-white text-indigo-900 font-bold px-6 py-3 rounded-xl hover:bg-gray-100 transition shadow-lg whitespace-nowrap">
                Ver Entradas
              </button>
            </div>

          </div>
        </div>
      </div>
    </div>
  );
};

export default Hero;